"use client";

import { useTheme } from "@/context/theme-provider";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { Moon, Sun } from "lucide-react";
import { TooltipWrapper } from "./TooltipWrapper";

type ThemeToggleProps = React.ComponentProps<typeof Button>

export function ThemeToggle({ className, ...props }: ThemeToggleProps) {
  const { theme, toggleTheme } = useTheme();

  const handleThemeToggle = (event: React.MouseEvent<HTMLButtonElement>) => {
    const { clientX: x, clientY: y } = event;
    toggleTheme({ x, y });
  };

  return (
    <TooltipWrapper label="Toggle theme" asChild>
      <Button className={cn("cursor-pointer", className)} {...props} onClick={handleThemeToggle}>
        {theme === "light" ? <Sun /> : <Moon />}
      </Button>
    </TooltipWrapper>
  );
}
