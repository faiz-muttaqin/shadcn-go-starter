import { createFileRoute } from '@tanstack/react-router'
import { useEffect, useState } from "react";
import { Header } from "@/features/home/header";
import { Hero } from "@/features/home/hero";
import { ThemePresetSelector } from "@/features/home/theme-preset-selector";

export const Route = createFileRoute('/')({
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
    <div className="bg-background text-foreground flex min-h-[100dvh] flex-col items-center justify-items-center">
      <Header
        isScrolled={isScrolled}
        mobileMenuOpen={mobileMenuOpen}
        setMobileMenuOpen={setMobileMenuOpen}
      />
      <main className="w-full flex-1">
        <Hero />
        <ThemePresetSelector />
        {/* <Testimonials />
        <Features />
        <HowItWorks />
        <Roadmap />
        <FAQ />
        <CTA /> */}
      </main>
      {/* <Footer /> */}
    </div>
  );
}
