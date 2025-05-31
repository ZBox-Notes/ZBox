import { AppSidebar } from "@/components/app-sidebar"
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar"
import { Outlet } from "react-router"
import "./index.css"

export default function Layout() {
    return (
        <SidebarProvider>
            <AppSidebar />
            <main className="w-full px-8 py-4 mt-4">
                <SidebarTrigger className="outline outline-solid outline-sidebar-border" />
                <Outlet />
            </main>
        </SidebarProvider>
    )
}