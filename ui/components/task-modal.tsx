import { Button } from "@/components/ui/button"

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Plus } from "lucide-react"

import { FormEvent, useContext, useState } from "react"
import { TasksContext, TasksContextData } from "@/contexts/tasks-context"
import { Label } from "./ui/label"
import { Textarea } from "./ui/textarea"

export function TaskModal() {
  const { createTask, isLoading } = useContext(TasksContext) as TasksContextData

  const [description, setDescription] = useState('')
  const [dialogOpen, setDialogOpen] = useState(false)


  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    try {
      await createTask(description)
      setDescription('')
      setDialogOpen(false)
    } catch {

    } finally {
    }
  }

  return (
    <Dialog open={dialogOpen} onOpenChange={open => setDialogOpen(open)}>
      <DialogTrigger asChild>
        <Button>
          <Plus />
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Nova tarefa</DialogTitle>
        </DialogHeader>
        <form className="flex flex-col gap-2" onSubmit={onSubmit}>
          <Label>Descrição</Label>
          <Textarea
            value={description}
            onChange={e => setDescription(e.target.value)}
            disabled={isLoading}
            className="min-h-[150px] max-h-[185px] " />
          <Button disabled={isLoading} type="submit">Salvar</Button>
        </form>
      </DialogContent>
    </Dialog>
  )
}
