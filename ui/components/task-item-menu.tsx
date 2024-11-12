"use client"

import {
  Check,
  Menu,
  RefreshCw,
  Trash,
  Undo
} from "lucide-react"

import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { TaskDeleteConfirm } from "./task-delete-confirmation"

import * as taskServices from '@/services/tasks-services'
import { useContext, useState } from "react"
import { TasksContext, TasksContextData } from "@/contexts/tasks-context"
import { useToast } from "@/hooks/use-toast"

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
  const { isLoading, deleteTask } = useContext(TasksContext) as TasksContextData
  const [deleteTaskDialogOpen, setDeleteTaskOpen] = useState(false)

  const { toast } = useToast()

  const callbackDeleteConfirm = () => {
    deleteTask(task.id)
      .then(() => {
        toast({
          title: "Tarefa removida",
        })
      })
      .catch((e: Error) => {
        toast({
          title: e.message,
          variant: 'destructive'
        })
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
          <DropdownMenuItem
            hidden={isLoading}
            onClick={() => setDeleteTaskOpen(true)}>
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