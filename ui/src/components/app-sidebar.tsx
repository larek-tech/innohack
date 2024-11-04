import { Calendar, ChartNoAxesCombined, MessagesSquare } from "lucide-react";
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import { Link, useLocation } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import ChatSessionService from "@/api/ChatSessionService";
import { SessionDto } from "@/api/models";
import { ChatSessionList } from "@/components/chatSessionList";

// Menu items
const items = [
  {
    title: "Дашборд",
    url: "/dash",
    icon: ChartNoAxesCombined,
  },
  {
    title: "Чат",
    url: "/chat",
    icon: MessagesSquare,
  }
];

export function AppSidebar() {
  const [sessions, setSessions] = useState<SessionDto[]>([]);
  const location = useLocation();
  const isChatPage = location.pathname.startsWith("/chat");

  useEffect(() => {
    if (isChatPage) {
      ChatSessionService.getSessions().then((response) => {
        setSessions(response);
      });
    }
  }, [isChatPage]);

  return (
    <Sidebar>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>misis banch space</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <Link to={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
        {isChatPage && <ChatSessionList sessions={sessions} />}
      </SidebarContent>
    </Sidebar>
  );
}