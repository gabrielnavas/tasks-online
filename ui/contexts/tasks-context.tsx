"use client"

import * as taskServices from "@/services/tasks-services"
import React, { useEffect, useState } from "react"
import { HttpInternalServerError } from "./exceptions"

type Task = {
  id: string
  description: string
  done: boolean
  createdAt: Date
  updatedAt?: Date | undefined
}

export type TasksContextData = {
  isLoading: boolean
  tasks: Task[]
  createTask: (description: string) => Promise<Task>
}

export const TasksContext = React.createContext<TasksContextData | null>(null)

type Props = {
  children: React.ReactNode
}

export const TasksProvider = ({ children }: Props) => {
  const [isLoading, setIsLoading] = useState(false)
  const [tasks, setTasks] = useState<Task[]>([])

  useEffect(() => {
    (async () => {
      try {
        setIsLoading(true)
        const tasks = await taskServices.findTasks({
          size: 10,
          page: 1,
          query: ''
        })
        setTasks(tasks)
      }
      catch (err) {
        console.log(err);
        throw new HttpInternalServerError()
      } finally {
        setIsLoading(false)
      }
    })()
  }, [])

  const createTask = async (description: string): Promise<Task> => {
    try {
      setIsLoading(true)
      const task = await taskServices.createTask({
        description,
      })
      setTasks(prev => [task, ...prev])
      return task
    } catch (err) {
      console.log(err);
      throw new HttpInternalServerError()
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <TasksContext.Provider value={{
      isLoading,
      tasks,
      createTask
    }}>
      {children}
    </TasksContext.Provider>
  )
}