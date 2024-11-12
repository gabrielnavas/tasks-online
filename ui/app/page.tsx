import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area";
import { TaskItem } from "../components/task-item";

export default function Home() {
  return (
    <section className="flex gap-4">
      <ScrollArea className="flex flex-nowrap">
        <div className="flex w-max space-x-4 p-4 gap-4">
          {
            new Array(20).fill('')
              .map((_, index) => (
                <TaskItem
                  key={index.toString()}
                  task={{
                    id: index.toString(),
                    content: generateLorem(),
                    createdAt: new Date(new Date().toUTCString()),
                    updatedAt: index % 2 === 0 ? new Date(new Date().toUTCString()) : undefined,
                    done: index % 2 === 0,
                  }} />
              ))
          }
          <ScrollBar orientation="horizontal" />
        </div>
      </ScrollArea>
    </section>
  );
}

function generateLorem() {
  return "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
}
