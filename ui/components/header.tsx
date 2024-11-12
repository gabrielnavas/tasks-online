import { Card, CardContent } from "@/components/ui/card"
import { ThemeMenu } from "./theme-menu"

export const Header = () => {
  return (
    <Card>
      <CardContent className="flex items-center justify-between p-2 px-6">
        <div className="font-semibold">Task Online</div>
        <ThemeMenu />
      </CardContent>
    </Card>
  )
}