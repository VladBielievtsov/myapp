import React from "react";
import LoginPage from "./LoginPage";
import { Metadata } from "next";

export default function page() {
  return <LoginPage />;
}

export const metadata: Metadata = {
  title: "Login",
};
