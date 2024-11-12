import { Card, CardContent, CardHeader } from "@/components/ui/card"
import { Separator } from "@/components/separator"
import { distanceToNow } from "@/lib/date"
import { Check, RefreshCw, ShieldCheck, X } from "lucide-react"
import { TaskItemMenu } from "./task-item-menu"

type Props = {
  task: {
    id: string
    description: string
    done: boolean
    createdAt: Date
    updatedAt?: Date | undefined
  }
}

export const TaskItem = ({ task }: Props) => {
  return (
    <Card className="h-[100%] w-[500px] p-4">
      <CardHeader className="flex">
        <div className="flex items-center justify-between">
          <div className={`flex flex-col gap-1 ${task.done && 'text-gray-400'}`}>
            {
              task.done ? (
                <span className="flex items-center gap-2 text-sm">
                  <Check color='green' size={15} />
                  Finalizado
                </span>
              ) : (
                <span className="flex items-center gap-2 text-sm">
                  <X color='red' size={15} />
                  Em andamento
                </span>
              )
            }
            {task.updatedAt && (
              <span className="flex items-center gap-2 text-sm">
                <RefreshCw size={12} />
                Atualizado {distanceToNow(task.updatedAt)}
              </span>
            )}
            <span className="flex items-center gap-2 text-sm">
              <ShieldCheck size={12} />
              Criado {distanceToNow(task.createdAt)}
            </span>
          </div>
          <div>
            <TaskItemMenu task={task} />
          </div>
        </div>
        <Separator className="my-4" />
      </CardHeader>
      <CardContent
        className={`break-words ${task.done && 'line-through text-gray-400'}`}>
        {task.description}
      </CardContent>
    </Card>
  )
}