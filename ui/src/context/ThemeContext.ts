import { createContext } from 'react';

// Theme types
export type ThemeType = 'light' | 'dark' | 'system';

export interface ThemeContextType {
  theme: ThemeType;
  setTheme: (theme: ThemeType) => void;
}

// Create the context with default values
export const ThemeContext = createContext<ThemeContextType>({
  theme: 'system',
  setTheme: (theme: ThemeType) => { /* Implementation provided by provider */ },
}); 