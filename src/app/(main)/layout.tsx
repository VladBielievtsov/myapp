import { Metadata } from "next";
import React from "react";
import App from "./App";
import NavBar from "./NavBar";

export default function ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <App>
      <main>
        <nav className="h-[60px] w-full">
          <NavBar />
        </nav>
        <section>
          <div className="container">{children}</div>
        </section>
      </main>
    </App>
  );
}

export const metadata: Metadata = {
  title: {
    template: "%s | Umami",
    default: "Umami",
  },
};
