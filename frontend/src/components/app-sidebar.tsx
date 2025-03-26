import logo from "@/assets/logo.svg"
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarHeader,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem
} from "@/components/ui/sidebar"
import { Archive, House, Settings, StickyNote, User } from "lucide-react"

export function AppSidebar() {
    return (
        <Sidebar className="items-start">
            <SidebarHeader className="items-start px-4 py-2">
                <div className="flex px-4 py-2 items-center justify-center rounded-(--radius-sm) bg-white w-full outline outline-solid outline-sidebar-border">
                    <img src={logo} width={50} height={50} />
                    <h2 className="text-4xl align-middle">Box</h2>
                </div>
            </SidebarHeader>
            <SidebarContent>
                <SidebarGroup className="items-start text-left">
                    <SidebarGroupContent>
                        <SidebarMenu>
                            <SidebarMenuItem key={"home"}>
                                <SidebarMenuButton asChild className="hover:bg-gray-100">
                                    <a href="/">
                                        <House />
                                        Home
                                    </a>
                                </SidebarMenuButton>
                            </SidebarMenuItem>
                            <SidebarMenuItem key={"notes"}>
                                <SidebarMenuButton asChild className="hover:bg-gray-100">
                                    <a href="/notes">
                                        <StickyNote />
                                        Notes
                                    </a>
                                </SidebarMenuButton>
                            </SidebarMenuItem>
                            <SidebarMenuItem key={"boxes"}>
                                <SidebarMenuButton asChild className="hover:bg-gray-100">
                                    <a href="/boxes">
                                        <Archive />
                                        Boxes
                                    </a>
                                </SidebarMenuButton>
                            </SidebarMenuItem>
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
                <SidebarGroup>
                    <SidebarGroupLabel>
                        <span className="text-lg">
                            Recent notes
                        </span>
                    </SidebarGroupLabel>
                </SidebarGroup>
            </SidebarContent>
            <SidebarFooter className="px-4 py-2">
                <div className="flex px-4 py-2 items-start justify-left rounded-(--radius-md) bg-white w-full outline outline-solid outline-sidebar-border">
                    <User strokeWidth={1} />
                    <h3 className="mx-auto">
                        David René
                    </h3>
                    <Settings strokeWidth={1} />
                </div>
            </SidebarFooter>
        </Sidebar>
    )
}