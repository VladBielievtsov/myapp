"use client";

import { cn } from "@/lib/utils";
import { CircleUserRound, Moon, Soup, Sun, User } from "lucide-react";
import Link from "next/link";
import React from "react";
import { usePathname } from "next/navigation";
import Button from "@/components/ui/Button";
import { useThemeStore } from "@/store";
import {
  DropdownMenu,
  DropdownContent,
  DropdownToggle,
} from "@/components/ui/DropdownMenu";

export default function NavBar() {
  const { theme, toggleTheme } = useThemeStore();
  const pathname = usePathname();

  const links = [
    { label: "Dashboard", url: "/dashboard" },
    { label: "Link", url: "/link" },
  ];

  const accLinks = [
    { icon: <User size={16} />, label: "Profile", url: "/profile" },
  ];

  return (
    <div className="grid grid-cols-nav items-center h-[60px] bg-zinc-50 border-b border-zinc-200 px-[20px] dark:bg-zinc-800 dark:border-zinc-600">
      <div className="flex items-center gap-[10px] text-base font-bold text-zinc-800 dark:text-zinc-300">
        <Soup size={21} strokeWidth={2} />
        <span>myapp</span>
      </div>
      <div className="flex items-center font-bold gap-[30px] px-[40px]">
        {links.map(({ url, label }) => (
          <Link
            href={url}
            key={url}
            className={cn(
              "text-zinc-800 border-b-2 border-transparent hover:border-blue-500 text-sm leading-[60px] dark:text-zinc-400",
              pathname.startsWith(url) && "border-blue-500 dark:text-zinc-300"
            )}
          >
            {label}
          </Link>
        ))}
      </div>
      <div className="flex items-center justify-end gap-2">
        <Button variant={["icon", "ghost"]} onCkick={toggleTheme}>
          {theme === "dark" ? <Moon size={21} /> : <Sun size={21} />}
        </Button>
        <DropdownMenu>
          <DropdownToggle>
            <Button variant={["icon", "ghost"]}>
              <CircleUserRound size={21} strokeWidth={2} />
            </Button>
          </DropdownToggle>
          <DropdownContent align="right" className="min-w-[198px] ">
            <div className="bg-zinc-50 border-b border-zinc-200 text-zinc-500 py-3 px-5 text-sm dark:bg-zinc-800 dark:border-zinc-600 dark:text-zinc-400">
              admin
            </div>
            <div>
              {accLinks.map(({ url, icon, label }) => (
                <Link
                  key={url}
                  href={url}
                  className="flex bg-white hover:bg-zinc-50 gap-3 items-center py-3 px-5 whitespace-nowrap text-sm cursor-pointer dark:bg-zinc-900 dark:hover:bg-zinc-800 text-zinc-800 dark:text-white"
                >
                  {icon}
                  {label}
                </Link>
              ))}
            </div>
          </DropdownContent>
        </DropdownMenu>
      </div>
    </div>
  );
}
