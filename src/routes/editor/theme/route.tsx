import { Header } from '@/features/home/header';
import { createFileRoute, Outlet } from '@tanstack/react-router'
import { useEffect, useState } from 'react';

export const Route = createFileRoute('/editor/theme')({
  component: RouteComponent,
})

function RouteComponent() {
    const [isScrolled, setIsScrolled] = useState(false);
    const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  
    useEffect(() => {
      const handleScroll = () => {
        if (window.scrollY > 10) {
          setIsScrolled(true);
        } else {
          setIsScrolled(false);
        }
      };
  
      window.addEventListener("scroll", handleScroll);
      return () => window.removeEventListener("scroll", handleScroll);
    }, []);
  
  return (
    <div className="relative isolate flex h-svh flex-col overflow-hidden">
      <Header
        isScrolled={isScrolled}
        mobileMenuOpen={mobileMenuOpen}
        setMobileMenuOpen={setMobileMenuOpen}
      />
      <main className="isolate flex flex-1 flex-col overflow-hidden">
        <Outlet />
      </main>
    </div>
  );
}
