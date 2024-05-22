import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  darkMode: "selector",
  theme: {
    container: {
      center: true,
      padding: "20px",
      screens: {
        DEFAULT: "1320px",
      },
    },
    extend: {
      gridTemplateColumns: {
        nav: "max-content 1fr 1fr",
      },
    },
  },
  plugins: [],
};
export default config;
