import { cn } from "@/lib/utils";
import React from "react";

interface ButtonProps {
  children: React.ReactNode;
  variant?: Array<"default" | "ghost" | "outline" | "icon">;
  className?: string;
  type?: "button" | "submit" | "reset";
  onCkick?: React.MouseEventHandler<HTMLButtonElement>;
}

export default function Button({
  children,
  className,
  type,
  variant = ["default"],
  onCkick,
}: ButtonProps) {
  return (
    <button
      type={type || "button"}
      className={cn(
        "text-base bg-zinc-700 text-white dark:text-zinc-900 dark:bg-white hover:bg-opacity-80 transition flex items-center justify-center border border-transparent rounded-lg cursor-pointer gap-2.5 min-h-[40px] px-4 whitespace-nowrap",
        variant.includes("icon") && "w-[50px] px-0",
        variant.includes("ghost") &&
          "bg-transparent dark:bg-transparent text-zinc-900 dark:text-white hover:bg-zinc-200 dark:hover:bg-zinc-600",
        variant.includes("outline") &&
          "bg-transparent dark:bg-transparent text-zinc-900 dark:text-white border-zinc-200 dark:border-zinc-600 dark:hover:bg-zinc-600 hover:bg-zinc-200",
        className
      )}
      onClick={onCkick}
    >
      {children}
    </button>
  );
}
