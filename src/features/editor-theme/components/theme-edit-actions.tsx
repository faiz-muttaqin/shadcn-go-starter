import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip";
import { useUpdateTheme } from "@/hooks/themes";
import { useEditorStore } from "@/stores/editor-store";
import { type Theme } from "@/types/theme";
import { Check, X } from "lucide-react";
import { useNavigate, useLocation } from '@tanstack/react-router'
import { useState } from "react";
import { ThemeSaveDialog } from "./theme-save-dialog";

interface ThemeEditActionsProps {
  theme: Theme;
  disabled?: boolean;
}

const ThemeEditActions: React.FC<ThemeEditActionsProps> = ({ theme, disabled = false }) => {
  const navigate = useNavigate();
  const location = useLocation();
  const updateThemeMutation = useUpdateTheme();
  const { themeState, applyThemePreset } = useEditorStore();
  const [isNameDialogOpen, setIsNameDialogOpen] = useState(false);

  // Preserve current search params when navigating back to the editor.
  // `location.search` from TanStack Router may be a parsed object containing
  // numbers or arrays; `navigate` expects string values. Normalize all values
  // to strings (join arrays) so the shape matches `Record<string, string>`.
  const mainEditorUrlSearch: Record<string, string> | undefined = (() => {
    const s = location.search;
    if (!s) return undefined;

    // If the router provided a raw query string, parse it directly.
    if (typeof s === "string") {
      return Object.fromEntries(new URLSearchParams(s));
    }

    // Otherwise we expect an object-like shape. Narrow it to a record of
    // unknown so we can inspect values safely without using `any`.
    if (typeof s === "object" && s !== null) {
      const obj = s as Record<string, unknown>;
      const entries: [string, string][] = [];
      for (const [k, v] of Object.entries(obj)) {
        if (v === undefined || v === null) continue;
        if (Array.isArray(v)) {
          entries.push([k, v.map((x) => String(x)).join(",")]);
        } else {
          entries.push([k, String(v)]);
        }
      }
      return Object.fromEntries(entries);
    }

    return undefined;
  })();

  const handleThemeEditCancel = () => {
    // Keep the current search params for tab persistence
    navigate({ to: '/editor/theme', search: mainEditorUrlSearch });
    applyThemePreset(themeState?.preset || "default");
  };

  const handleSaveTheme = async (newName: string) => {
    const dataToUpdate: {
      id: string;
      name?: string;
      styles?: Theme["styles"];
    } = {
      id: theme.id,
    };

    if (newName !== theme.name) {
      dataToUpdate.name = newName;
    } else {
      dataToUpdate.name = theme.name;
    }

    if (themeState.styles) {
      dataToUpdate.styles = themeState.styles;
    }

    if (!dataToUpdate.name && !dataToUpdate.styles) {
      setIsNameDialogOpen(false);
      return;
    }

    try {
      const result = (await updateThemeMutation.mutateAsync(dataToUpdate)) as Theme | undefined;
      if (result) {
        setIsNameDialogOpen(false);
        navigate({ to: '/editor/theme', search: mainEditorUrlSearch });
        applyThemePreset(result.id || themeState?.preset || "default");
      }
    } catch (error) {
      console.error("Failed to update theme:", error);
    }
  };

  const handleThemeEditSave = () => {
    setIsNameDialogOpen(true);
  };

  return (
    <>
      <div className="bg-card/80 text-card-foreground flex items-center">
        <div className="flex min-h-14 flex-1 items-center gap-2 px-4">
          <div className="flex animate-pulse items-center gap-2">
            <div className="h-2 w-2 rounded-full bg-blue-500" />
            <span className="text-card-foreground/60 text-sm font-medium">Editing</span>
          </div>
          <span className="max-w-56 truncate px-2 text-sm font-semibold">{theme.name}</span>
        </div>

        <Separator orientation="vertical" className="bg-border h-8" />

        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="size-14 shrink-0 rounded-none"
                onClick={handleThemeEditCancel}
                disabled={disabled}
              >
                <X className="h-4 w-4" />
              </Button>
            </TooltipTrigger>
            <TooltipContent>Cancel changes</TooltipContent>
          </Tooltip>
        </TooltipProvider>

        <Separator orientation="vertical" className="bg-border h-8" />

        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="size-14 shrink-0 rounded-none"
                onClick={handleThemeEditSave}
                disabled={disabled}
              >
                <Check className="h-4 w-4" />
              </Button>
            </TooltipTrigger>
            <TooltipContent>Save changes</TooltipContent>
          </Tooltip>
        </TooltipProvider>
      </div>

      <ThemeSaveDialog
        open={isNameDialogOpen}
        onOpenChange={setIsNameDialogOpen}
        onSave={handleSaveTheme}
        isSaving={updateThemeMutation.isPending}
        initialThemeName={theme.name}
        title="Save Theme Changes"
        description="Confirm or update the theme name before saving."
        ctaLabel="Save Changes"
      />
    </>
  );
};

export default ThemeEditActions;
