"use client";

import React from "react";
import { Inter } from "next/font/google";
import { cn } from "@/lib/utils";
import { useThemeStore } from "@/store";

const inter = Inter({ subsets: ["latin"] });

export default function ThemeProvider({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const { theme } = useThemeStore();
  return (
    <html lang="en" data-scroll="0" className={theme}>
      <head>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta name="robots" content="noindex,nofollow" />
      </head>
      <body className={cn(inter.className, "dark:bg-zinc-900 bg-white")}>
        {children}
      </body>
    </html>
  );
}
