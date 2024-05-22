import type { Metadata } from "next";

import "./globals.css";
import ThemeProvider from "./ThemeProvider";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return <ThemeProvider>{children}</ThemeProvider>;
}

export const metadata: Metadata = {
  title: {
    template: "%s | My App",
    default: "My App",
  },
};
