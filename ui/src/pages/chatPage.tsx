import { useState, useRef, useEffect } from 'react';
import { Send } from 'lucide-react';
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ScrollArea } from "@/components/ui/scroll-area";
import { observer } from 'mobx-react-lite';
import { useToast } from '@/hooks/use-toast';
import { WS_URL } from '@/config';
import { Route, useNavigate, useParams } from '@tanstack/react-router';
import { LOCAL_STORAGE_KEY } from '@/auth/AuthProvider';
import ChatSessionService from '@/api/ChatSessionService';
import { SessionDto, QueryDto, ResponseDto, SessionContentDto } from '@/api/models';
import Markdown from 'react-markdown';


interface ChatMessage {
    data: ResponseDto;
    sender: 'user' | 'chat';
    // graph data for line chart
    graphData?: any;
}

function mapSessionContentDtoToMessages(data: SessionContentDto[]): ChatMessage[] {
    return data.flatMap(item => [
        {
            data: {
                queryId: item.query.id,
                sources: [],
                filenames: [],
                charts: [],
                description: item.query.prompt,
                multipliers: [],
                createdAt: item.query.created_at,
                error: "",
                isLast: false,
            },
            sender: "user"
        },
        item.response && {
            data: item.response,
            sender: "chat"
        }
    ].filter(Boolean));
}

const chatMessage = (msg: ChatMessage, index: number) => (
    <div
        key={index}
        className={`flex ${msg.sender === 'user' ? 'justify-end' : 'justify-start'}`}
    >
        <div className={`max-w-[70%] p-3 rounded-lg ${msg.sender === 'user' ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-800'}`}>
            <Markdown>{msg.data.description}</Markdown>
        </div>
    </div>
);

const ChatInterface = observer(() => {
    const navigate = useNavigate();
    const { toast } = useToast();
    const [sessionId, setSessionId] = useState<string | null>(null);
    const [sessions, setSessions] = useState<SessionDto[]>([]);
    const [messages, setMessages] = useState<ChatMessage[]>([]);
    const [inputMessage, setInputMessage] = useState('');
    const [socket, setSocket] = useState<WebSocket | null>(null);
    const scrollAreaRef = useRef<HTMLDivElement>(null);

    // Scroll to the bottom when messages change
    useEffect(() => {
        if (scrollAreaRef.current) {
            scrollAreaRef.current.scrollTop = scrollAreaRef.current.scrollHeight;
        }
    }, [messages]);

    useEffect(() => {
        // fetch session from server if they are empty send request to create a new session
        if (sessions.length === 0) {
            ChatSessionService.getSessions().then((response) => {
                setSessions(response);
                if (response.length === 0) {
                    ChatSessionService.createSession().then((res) => {
                        const newChatId = res.id;
                        toast({
                            title: 'Chat Created',
                            description: `Chat with ID ${newChatId} created.`,
                        });
                        navigate({ to: `/chat?sessionId=${newChatId}` });
                        window.location.reload();
                    }).catch((err) => {
                        toast({
                            title: 'Error',
                            description: err.message,
                        });
                        navigate({ to: '/' });
                    });
                } else {
                    setSessionId(response[0].id);
                }
            });
        }
    }, []);
    // Load initial messages from the server
    useEffect(() => {
        if (sessionId != null) {
            ChatSessionService.getSessionContent(sessionId).then((res) => {
                const initialMessages = mapSessionContentDtoToMessages(res);
                setMessages(initialMessages);
            }).catch((err) => {
                toast({
                    title: 'Error',
                    description: err.message,
                });
                ChatSessionService.createSession().then((res) => {
                    const newChatId = res.id;
                    toast({
                        title: 'Chat Created',
                        description: `Chat with ID ${newChatId} created.`,
                    });
                    navigate({ to: `/chat/${newChatId}` });
                    window.location.reload();
                }).catch((err) => {
                    toast({
                        title: 'Error',
                        description: err.message,
                    });
                    navigate({ to: '/' });
                });
            });
        }
    }, [sessionId]);

    // Initialize WebSocket and handle incoming messages
    useEffect(() => {
        if (sessionId != null) {
            const ws = new WebSocket(`${WS_URL}/${sessionId}`);
            setSocket(ws);

            const req: QueryDto = {
                id: 0,
                prompt: `${JSON.parse(localStorage.getItem(LOCAL_STORAGE_KEY) as string)?.user?.token}`,
                createdAt: null,
            };

            ws.addEventListener("open", () => {
                ws.send(JSON.stringify(req));
            });

            ws.addEventListener("message", (event) => {
                const response: ResponseDto = JSON.parse(event.data);

                setMessages((prevMessages) => {
                    const lastMessage = prevMessages[prevMessages.length - 1];

                    if (response.isLast) {
                        return [...prevMessages.slice(0, -1), { data: response, sender: "chat" }];
                    }
                    if (lastMessage && response.queryId === lastMessage.data.queryId && lastMessage.data.description) {
                        const updatedMessage = {
                            ...lastMessage,
                            data: {
                                ...lastMessage.data,
                                description: lastMessage.data.description + response.description,
                            }
                        };
                        return [...prevMessages.slice(0, -1), updatedMessage];
                    } else {
                        return [...prevMessages, { data: response, sender: "chat" }];
                    }
                });
            });

            return () => {
                ws.close();
            };
        }
    }, [sessionId]);

    const handleSendMessage = () => {
        const req: QueryDto = {
            id: 0,
            prompt: inputMessage,
            createdAt: null,
        };

        if (socket) {
            socket.send(JSON.stringify(req));
            setMessages((prevMessages) => [
                ...prevMessages,
                {
                    data: {
                        queryId: req.id,
                        sources: [],
                        filenames: [],
                        charts: [],
                        description: req.prompt,
                        multipliers: [],
                        createdAt: new Date(),
                        error: "",
                        isLast: false,
                    },
                    sender: "user",
                }
            ]);
            setInputMessage('');
        }
    };

    return (
        <div className="flex flex-col flex-grow w-full h-full">
            <ScrollArea className="flex-grow p-4 space-y-4" ref={scrollAreaRef}>
                {messages.map((message, index) => chatMessage(message, index))}
                <div />
            </ScrollArea>
            <div className="p-4 border-t">
                <div className="flex space-x-2">
                    <Input
                        type="text"
                        placeholder="Введите сообщение..."
                        value={inputMessage}
                        onChange={(e) => setInputMessage(e.target.value)}
                        onKeyPress={(e) => {
                            if (e.key === 'Enter') {
                                handleSendMessage();
                            }
                        }}
                        className="flex-grow"
                    />
                    <Button onClick={handleSendMessage}>
                        <Send className="w-4 h-4" />
                    </Button>
                </div>
            </div>
        </div>
    );
});

export default ChatInterface;