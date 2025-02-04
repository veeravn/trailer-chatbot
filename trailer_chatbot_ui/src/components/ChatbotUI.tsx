import { useEffect, useState } from "react";
import { Button } from "./Button";
import { Input } from "./Input";
import { Card, CardContent } from "./Card";

const API_URL = "http://localhost:3000/chat"; // API endpoint

type ChatMessage = {
  role: "User" | "Bot";
  content: string;
};

export default function ChatbotUI() {
  const [message, setMessage] = useState("");
  const [chatHistory, setChatHistory] = useState<ChatMessage[]>([]);
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [isChatActive, setIsChatActive] = useState(false);

  useEffect(() => {
    if (isChatActive) {
      setChatHistory([{ role: "Bot", content: "Hello! How can I assist you today?" }]);
      const ws = new WebSocket("ws://localhost:3000/ws");
      
      ws.onopen = () => console.log("Connected to WebSocket server");
      ws.onmessage = (event) => setChatHistory((prev) => [...prev, { role: "Bot", content: event.data }]);
      ws.onerror = (error) => console.error("WebSocket error:", error);
      ws.onclose = () => console.log("WebSocket connection closed");

      setSocket(ws);

      return () => ws.close();
    }
  }, [isChatActive]);

  const sendMessage = async () => {
    if (!message.trim()) return;

    setChatHistory((prev) => [...prev, { role: "User", content: message }]);
    setMessage("");

    try {
      const response = await fetch(API_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ message }),
      });
      const data = await response.json();
      setChatHistory((prev) => [...prev, { role: "Bot", content: data.response }]);
    } catch (error) {
      console.error("Error sending message:", error);
    }
  };

  return (
    <div className="fixed bottom-4 right-4">
      <Button onClick={() => setIsChatActive(true)} className="bg-blue-600 text-white px-4 py-2 rounded-full shadow-lg">
        💬 Chat
      </Button>
      {isChatActive && (
        <div className="fixed bottom-16 right-4 bg-white shadow-lg rounded-lg w-80 h-96 p-4 border flex flex-col">
          <div className="flex justify-between items-center border-b pb-2">
            <h2 className="text-lg font-semibold">Chatbot</h2>
            <Button className="text-red-500" onClick={() => setIsChatActive(false)}>X</Button>
          </div>
          <Card className="flex-grow overflow-y-auto p-2">
            <CardContent>
              {chatHistory.map((chat, index) => (
                <div key={index} className={`p-2 ${chat.role === "User" ? "text-right" : "text-left"}`}>
                  <strong>{chat.role}:</strong> {chat.content}
                </div>
              ))}
            </CardContent>
          </Card>
          <div className="flex gap-2 mt-2">
            <Input value={message} onChange={(e) => setMessage(e.target.value)} placeholder="Type a message..." className="flex-1 border p-2 rounded-md" />
            <Button onClick={sendMessage} className="bg-blue-600 text-white px-4 py-2 rounded-md">Send</Button>
          </div>
        </div>
      )}
    </div>
  );
}