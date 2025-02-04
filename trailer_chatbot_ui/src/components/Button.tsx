export function Button({ children, onClick, className }: { children: React.ReactNode; onClick?: () => void; className?: string }) {
  return (
    <button onClick={onClick} className={`px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 ${className}`}>
      {children}
    </button>
  );
}
