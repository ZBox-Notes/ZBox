import { AppSidebar } from "@/components/app-sidebar"
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar"
import { Outlet } from "react-router"
import "./index.css"

export default function Layout() {
    return (
        <SidebarProvider>
            <AppSidebar />
            <SidebarTrigger className="my-4 mx-4" />
            <main className="w-full px-8 py-4">
                <Outlet />
            </main>
        </SidebarProvider>
    )
}