export function Input({ value, onChange, placeholder, className }: { value: string; onChange: (e: React.ChangeEvent<HTMLInputElement>) => void; placeholder?: string; className?: string }) {
  return <input type="text" value={value} onChange={onChange} placeholder={placeholder} className={`border p-2 rounded-md ${className}`} />;
}

