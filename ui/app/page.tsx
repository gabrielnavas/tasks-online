"use client"

import { TaskList } from "@/components/task-list";
import { TaskListHeader } from "@/components/task-list-header";
import { TasksProvider } from "@/contexts/tasks-context";

export default function Home() {

  return (
    <TasksProvider>
      <section className="light:bg-slate-100 flex flex-col gap-5 pt-4">
        <div className="flex justify-center">
          <TaskListHeader />
        </div>
        <TaskList />
      </section>
    </TasksProvider>
  );
}
