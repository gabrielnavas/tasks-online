"use client"

import { Check, Menu, RefreshCw, RotateCw, Trash, Undo, X } from "lucide-react"

import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

type Props = {
  task: {
    id: string
    content: string
    done: boolean
    createdAt: Date
    updatedAt?: Date | undefined
  }
}

export const TaskItemMenu = ({ task }: Props) => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="outline" size="icon">
          <Menu />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuItem onClick={() => { }}>
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
  )
}