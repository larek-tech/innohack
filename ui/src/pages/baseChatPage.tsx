import React from 'react';
import { useNavigate, useParams } from '@tanstack/react-router';
import axios from 'axios';
import { AppSidebar } from '@/components/app-sidebar';

const BaseChatPage: React.FC = () => {
  const navigate = useNavigate();
  const { sessionId } = useParams({strict: false});
  console.log(sessionId);
  const createNewChat = async () => {
    try {
      const response = await axios.post('/api/chats');
      const newChatId = response.data.id;
      navigate({ to: `/chat/${newChatId}` });
    } catch (error) {
      console.error('Failed to create new chat', error);
    }
  };

  return (
    <div className="flex h-screen">
      <AppSidebar />
      <div className="flex flex-col flex-grow p-4">
        <h1>Base Chat Page</h1>
        <button onClick={createNewChat}>Create New Chat</button>
      </div>
    </div>
  );
};

export default BaseChatPage;