export type CreateTaskDto = {
  description: string
}

export type UpdateTaskDto = {
  description: string
  done: boolean
}

export type FindTasksDto = {
  page: number
  size: number
  query: string
}

export type TaskDto = {
  id: string
  description: string
  done: boolean
  createdAt: Date
  updatedAt?: Date | undefined
}