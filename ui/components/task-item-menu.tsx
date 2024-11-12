"use client"

import { Check, Menu, RefreshCw, RotateCw, Trash, Undo, X } from "lucide-react"

import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { TaskDeleteConfirm } from "./task-delete-confirmation"

import * as taskServices from '@/services/tasks-services'
import { useState } from "react"

type Props = {
  task: {
    id: string
    description: string
    done: boolean
    createdAt: Date
    updatedAt?: Date | undefined
  }
}

export const TaskItemMenu = ({ task }: Props) => {
  const [deleteTaskDialogOpen, setDeleteTaskOpen] = useState(false)

  const callbackDeleteConfirm = () => {
    taskServices.deleteTask(task.id)
      .then(() => {
        console.log('Tarefa deletada!')
      })
      .catch(e => {
        alert(e)
      })
  }

  return (
    <>
      
      {/* modal para confirmação para deletar a task */}
      <TaskDeleteConfirm
        open={deleteTaskDialogOpen}
        setOpen={setDeleteTaskOpen}
        callbackConfirm={callbackDeleteConfirm}
      />

      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="outline" size="icon">
            <Menu />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuItem onClick={() => setDeleteTaskOpen(true)}>
            <Trash color='red' />
            Remover

          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => { }}>
            <RefreshCw color='yellow' />
            Atualizar
          </DropdownMenuItem>

          <DropdownMenuItem onClick={() => { }}>
            {
              task.done ? (
                <span className="flex gap-2">
                  <Undo color='red' />
                  Retornar
                </span>
              ) : (
                <span className="flex gap-2">
                  <Check color='green' />
                  Finalizar
                </span>
              )
            }
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  )
}