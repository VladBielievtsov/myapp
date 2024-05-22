import { cn } from "@/lib/utils";
import React, { forwardRef, InputHTMLAttributes } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  className?: string;
}

const Input = forwardRef<HTMLInputElement, InputProps>((props, ref) => {
  const { className, ...rest } = props;
  return (
    <input
      {...rest}
      ref={ref}
      className={cn(
        "w-full border min-h-[42px] px-5 text-zinc-800 text-base rounded border-zinc-200 dark:border-zinc-400 dark:bg-zinc-800 dark:text-zinc-200",
        className
      )}
    />
  );
});

export default Input;
