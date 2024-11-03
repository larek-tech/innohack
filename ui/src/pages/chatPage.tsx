
import { useState, useRef, useEffect } from 'react'
import { Send } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { ScrollArea } from "@/components/ui/scroll-area"
import { observer } from 'mobx-react-lite';
import { useStores } from '@/hooks/use-stores'
import { useToast } from '@/hooks/use-toast'
import { WS_URL } from '@/config'
import { useParams, useSearch } from '@tanstack/react-router'
import { LOCAL_STORAGE_KEY } from '@/auth/AuthProvider'
import ChatSessionService from '@/api/ChatSessionService'
import { SessionDto, QueryDto, ResponseDto } from '@/api/models'

type Message = {
    id: number
    content: string
    sender: 'user' | 'bot'
}

const ChatInterface = observer(() => {
    const search = useSearch({
        strict: false,
    })

    // const { rootStore } = useStores();
    const { toast } = useToast();
    const [session, setSession] = useState<SessionDto | null>(null)
    const [socket, setSocket] = useState<WebSocket | null>(null)
    const [messages, setMessages] = useState<ResponseDto[]>([])
    const [inputMessage, setInputMessage] = useState('')
    const scrollAreaRef = useRef<HTMLDivElement>(null)

    useEffect(() => {
        if (scrollAreaRef.current) {
            scrollAreaRef.current.scrollTop = scrollAreaRef.current.scrollHeight
        }
    }, [messages])



    useEffect(() => {
        if (search.chatId) {
            const ws = new WebSocket(`${WS_URL}/${search.chatId}`);
            setSocket(ws)
            const req: QueryDto = {
                id: 0,
                prompt: `${JSON.parse(localStorage.getItem(LOCAL_STORAGE_KEY) as string)?.user?.token}`,
                createdAt: null,
            }
            ws.addEventListener("open", (event) => {
                console.log(event)
                ws.send(JSON.stringify(req))
            });

            ws.addEventListener("message", (event) => {
                const response: ResponseDto = JSON.parse(event.data)
                if (response.queryId != messages[-1].queryId) {
                    setMessages([...messages, response])
                }
                if (response.queryId == messages[-1].queryId) {
                    const lastMessage = messages[-1]
                    lastMessage.description += response.description
                    setMessages([...messages, lastMessage])
                }

            })

        }
    }, [])

    const handleSendMessage = () => {
        const req: QueryDto = {
            id: 0,
            prompt: inputMessage,
            createdAt: null,
        }

    }

    return (
        <div className="flex flex-col h-screen max-w-2xl mx-auto">
            <ScrollArea className="flex-grow p-4 space-y-4" ref={scrollAreaRef}>
                {messages.map(message => (
                    <div
                        key={message.queryId}
                        className={`flex ${message.sender === 'user' ? 'justify-end' : 'justify-start'}`}
                    >
                        <div
                            className={`max - w - [70 %] p - 3 rounded - lg ${message.sender === 'user'
                                ? 'bg-blue-500 text-white'
                                : 'bg-gray-200 text-gray-800'
                                }`}
                        >
                            {message.description}
                        </div>
                    </div>
                ))}
            </ScrollArea>
            <div className="p-4 border-t">
                <div className="flex space-x-2">
                    <Input
                        type="text"
                        placeholder="Type your message..."
                        value={inputMessage}
                        onChange={(e) => setInputMessage(e.target.value)}
                        onKeyPress={(e) => {
                            if (e.key === 'Enter') {
                                handleSendMessage()
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
    )
})

export default ChatInterface;