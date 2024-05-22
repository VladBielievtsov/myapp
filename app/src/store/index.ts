import { getCookie, hasCookie } from "cookies-next";
import { create } from "zustand";

type Theme = "light" | "dark";

type ThemeStore = {
  theme: Theme;
  toggleTheme: () => void;
};

const getInitialTheme = (): Theme => {
  const storedTheme = "dark";
  return storedTheme === "dark" ? "dark" : "light";
};

export const useThemeStore = create<ThemeStore>((set) => {
  const initialTheme = getInitialTheme();

  return {
    theme: initialTheme,
    toggleTheme: () =>
      set((state) => {
        const newTheme = state.theme === "light" ? "dark" : "light";
        return { theme: newTheme };
      }),
  };
});
