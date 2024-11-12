import { Plus, Search } from "lucide-react"
import { Button } from "./ui/button"
import { Card, CardContent } from "./ui/card"
import { Input } from "./ui/input"
import { TaskModal } from "./task-modal"

export const TaskListHeader = () => {
  return (
    <Card>
      <CardContent className="flex items-center justify-between gap-4 p-4">
        <div className="flex w-full gap-2">
          <Input placeholder="Ex. Lavar o carro..." className="w-[500px]" />
          <Button variant='outline' className="w-[100px]">
            <Search />
          </Button>
        </div>
        <TaskModal />
      </CardContent>
    </Card>
  )
}