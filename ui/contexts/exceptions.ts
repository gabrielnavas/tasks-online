export class HttpInternalServerError extends Error {
  constructor(message: string = 'Ocorreu um problema. Tente novamente mais tarde') {
    super(message)
  }

}