import React from "react";

type Message = {
    id: string;
    role: "user" | "assistant";
    text: string;
};

export const ChatInterface: React.FC = () => {
    const [messages, setMessages] = React.useState<Message[]>([]);
    const [input, setInput] = React.useState("");
    const [isSending, setIsSending] = React.useState(false);
    const listRef = React.useRef<HTMLDivElement | null>(null);

    const scrollToBottom = () => {
        if (listRef.current) listRef.current.scrollTop = listRef.current.scrollHeight;
    };

    React.useEffect(() => {
        scrollToBottom();
    }, [messages]);

    const sendMessage = (text?: string) => {
        const t = (text ?? input).trim();
        if (!t) return;

        const userMsg: Message = { id: String(Date.now()) + "-u", role: "user", text: t };
        setMessages((m) => [...m, userMsg]);
        setInput("");

        // Simulate assistant reply
        setIsSending(true);
        window.setTimeout(() => {
            const reply: Message = {
                id: String(Date.now()) + "-a",
                role: "assistant",
                text: `Simulated reply: ${t}`,
            };
            setMessages((m) => [...m, reply]);
            setIsSending(false);
        }, 700);
    };

    const handleKeyDown: React.KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
        if (e.key === "Enter" && !e.shiftKey) {
            e.preventDefault();
            sendMessage();
        }
    };

    return (
        <div className="flex h-full max-h-[60vh] w-full max-w-2xl flex-col rounded-md border bg-background">
            <div className="px-4 py-3 border-b">
                <h3 className="text-sm font-semibold">AI Chat (local only)</h3>
                <p className="text-xs text-muted-foreground">Frontend-only demo chat box</p>
            </div>

            <div ref={listRef} className="flex-1 overflow-auto p-4 space-y-3">
                {messages.length === 0 ? (
                    <div className="text-sm text-muted-foreground">No messages yet — say hello.</div>
                ) : (
                    messages.map((m) => (
                        <div
                            key={m.id}
                            className={`max-w-[85%] p-2 rounded-md ${
                                m.role === "user" ? "ml-auto bg-primary/10 text-primary" : "bg-border/30"
                            }`}
                        >
                            <div className="whitespace-pre-wrap text-sm">{m.text}</div>
                        </div>
                    ))
                )}
            </div>

            <div className="border-t p-3">
                <div className="flex gap-2">
                    <textarea
                        value={input}
                        onChange={(e) => setInput(e.target.value)}
                        onKeyDown={handleKeyDown}
                        placeholder="Type a message and press Enter to send"
                        className="flex-1 resize-none rounded-md border px-3 py-2 text-sm"
                        rows={2}
                    />

                    <button
                        onClick={() => sendMessage()}
                        disabled={isSending || input.trim() === ""}
                        className="ml-2 inline-flex items-center rounded-md bg-primary px-3 py-2 text-sm text-white disabled:opacity-50"
                    >
                        {isSending ? "Sending…" : "Send"}
                    </button>
                </div>
            </div>
        </div>
    );
};

export default ChatInterface;
