import { useContext } from "react"
import { TaskItem } from "./task-item"
import { CircleX } from "lucide-react"
import { ScrollArea, ScrollBar } from "./ui/scroll-area"
import { TasksContext, TasksContextData } from "@/contexts/tasks-context"

export const TaskList = () => {
  const {
    isLoading,
    tasks
  } = useContext(TasksContext) as TasksContextData

  if (isLoading) {
    <div>
      <CircleX />
      carregando....
    </div>
  }

  if (tasks.length === 0) {
    return (
      <div className="flex justify-center items-center min-h-[120px]">
        <span className="font-semibol">
          Nenhuma tarefa listada.
        </span>
      </div>
    )
  }

  return (
    <ScrollArea className="flex flex-nowrap w-full">
      <div className="flex min-h-[450px] space-x-4 p-4 gap-4">
        {tasks.map(task => (
          <TaskItem
            key={task.id}
            task={task} />
        ))}
        <ScrollBar orientation="horizontal" />
      </div>
    </ScrollArea>

  )
}