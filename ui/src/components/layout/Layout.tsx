import React from 'react';
import { ThemeToggle } from '../ThemeToggle';

interface LayoutProps {
  children: React.ReactNode;
}

export function Layout({ children }: LayoutProps) {
  return (
    <div className="min-h-screen bg-background">
      <header className="sticky top-0 z-40 border-b bg-background ">
        <div className="flex h-16 items-center justify-between py-4">
          <div className="flex-1 flex items-center gap-2">
            <h1 className="text-xl font-bold ml-4">🚀 GOAD Dashboard</h1>
          </div>
          <div className="flex items-center gap-4 mr-4">
            <ThemeToggle />
          </div>
        </div>
      </header>
      <main className="container py-6 mx-auto flex flex-col items-center justify-center">{children}</main>
    </div>
  );
} 