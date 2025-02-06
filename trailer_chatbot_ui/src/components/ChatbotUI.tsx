import { useEffect, useState } from "react";
import { Button } from "./Button";
import { Input } from "./Input";
import { Card, CardContent } from "./Card";
import {
  BarChart, Bar, XAxis, YAxis, Tooltip, Legend, ResponsiveContainer,
} from "recharts";

const DASHBOARD_API = "http://localhost:3000/dashboard";
const CHAT_API = "http://localhost:3000/chat";

type ChatMessage = {
  role: "User" | "Bot";
  content: string;
};

export default function ChatbotUI() {
  const [message, setMessage] = useState("");
  const [chatHistory, setChatHistory] = useState<ChatMessage[]>([]);
  const [isChatActive, setIsChatActive] = useState(false);
  const [stats, setStats] = useState({ total: 0, completed: 0, pending: 0 });

  useEffect(() => {
    fetch(DASHBOARD_API)
      .then((res) => res.json())
      .then((data) => setStats(data))
      .catch((error) => console.error("Error fetching dashboard data:", error));
  }, []);

  const chartData: { name: string; value: number }[] = [
    { name: "Total Trailers", value: stats.total },
    { name: "Completed", value: stats.completed },
    { name: "Pending", value: stats.pending },
  ];

  const sendMessage = async () => {
    if (!message.trim()) return;

    setChatHistory((prev) => [...prev, { role: "User", content: message }]);
    setMessage("");

    try {
      const response = await fetch(CHAT_API, {
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
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-100 p-4">
      <h1 className="text-3xl font-bold mb-4">📊 Trailer Unloading Dashboard</h1>

      <div className="bg-white shadow-md rounded-lg p-6 w-full max-w-lg">
        <p className="text-lg">🚛 Total Trailers: <strong>{stats.total}</strong></p>
        <p className="text-lg text-green-600">✅ Completed: <strong>{stats.completed}</strong></p>
        <p className="text-lg text-red-600">⏳ Pending: <strong>{stats.pending}</strong></p>
      </div>

      <div className="w-full max-w-lg mt-6">
        <ResponsiveContainer width="100%" height={300}>
          <BarChart data={chartData}>
            <XAxis dataKey="name" />
            <YAxis />
            <Tooltip />
            <Legend />
            <Bar dataKey="value" fill="#8884d8" />
          </BarChart>
        </ResponsiveContainer>
      </div>

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
              <Input
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                placeholder="Type a message..."
                className="flex-1 border p-2 rounded-md"
              />
              <Button onClick={sendMessage} className="bg-blue-600 text-white px-4 py-2 rounded-md">Send</Button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
