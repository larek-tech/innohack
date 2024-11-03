import {
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import { Link } from "@tanstack/react-router";
import { SessionDto } from "@/api/models";

interface ChatSessionListProps {
  sessions: SessionDto[];
}

export function ChatSessionList({ sessions }: ChatSessionListProps) {
  return (
    <SidebarGroup>
      <SidebarGroupLabel>Chat Sessions</SidebarGroupLabel>
      <SidebarGroupContent>
        <SidebarMenu>
          {sessions.map((session) => (
            <SidebarMenuItem key={session.id}>
              <SidebarMenuButton asChild>
                <Link to={`/chat/${session.id}`}>
                  <span>{session.title}</span>
                </Link>
              </SidebarMenuButton>
            </SidebarMenuItem>
          ))}
        </SidebarMenu>
      </SidebarGroupContent>
    </SidebarGroup>
  );
}