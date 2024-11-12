import { CreateTaskDto, FindTasksDto, TaskDto, UpdateTaskDto } from "./dtos";

const url = `${process.env.NEXT_PUBLIC_URL_API}/tasks`

export const createTask = async (dto: CreateTaskDto): Promise<TaskDto> => {
  const response = await fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'applications/json',
    },
    body: JSON.stringify(dto)
  })

  if (response.status >= 400) {
    throw new Error('Houve um problema. Tente novamente mais tarde.')
  }

  const data = await response.json()
  return mapFromBodyToDto(data.task)
}


export const updateTask = async (taskId: string, dto: UpdateTaskDto): Promise<void> => {
  const response = await fetch(`${url}/${taskId}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'applications/json',
    },
    body: JSON.stringify(dto)
  })

  if (response.status >= 400) {
    throw new Error('Houve um problema. Tente novamente mais tarde.')
  }
}

export const deleteTask = async (taskId: string): Promise<void> => {
  const response = await fetch(`${url}/${taskId}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'applications/json',
    },
  })

  if (response.status >= 400) {
    throw new Error('Houve um problema. Tente novamente mais tarde.')
  }
}

export const findTasks = async ({ size, page, query }: FindTasksDto): Promise<TaskDto[]> => {
  const response = await fetch(`${url}?page=${page}&size=${size}&query=${query}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'applications/json',
    },
  })

  if (response.status >= 400) {
    throw new Error('Houve um problema. Tente novamente mais tarde.')
  }

  const data = await response.json()
  return data.tasks.map(mapFromBodyToDto)
}

const mapFromBodyToDto = (task: any) => {
  return {
    id: task.id,
    description: task.description,
    createdAt: new Date(task.createdAt),
    done: task.done,
  }
}