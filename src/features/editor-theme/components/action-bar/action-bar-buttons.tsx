import { Separator } from "@/components/ui/separator";
// import { useAIThemeGenerationCore } from "@/hooks/use-ai-theme-generation-core";
import { useEditorStore } from "@/stores/editor-store";
import { useThemePresetStore } from "@/stores/theme-preset-store";
import { MoreOptions } from "./components/more-options";
import { ThemeToggle } from "./components/theme-toggle";
import { EditButton } from "./components/edit-button";
import { ImportButton } from "./components/import-button";
import { ResetButton } from "./components/reset-button";
import { UndoRedoButtons } from "./components/undo-redo-buttons";
import { ShareButton } from "./components/share-button";
import { SaveButton } from "./components/save-button";
import { CodeButton } from "./components/code-button";

interface ActionBarButtonsProps {
  onImportClick: () => void;
  onCodeClick: () => void;
  onSaveClick: () => void;
  onShareClick: (id?: string) => void;
  isSaving: boolean;
}

export function ActionBarButtons({
  onImportClick,
  onCodeClick,
  onSaveClick,
  onShareClick,
  isSaving,
}: ActionBarButtonsProps) {
  const { themeState, resetToCurrentPreset, hasUnsavedChanges } = useEditorStore();
  // const { isGeneratingTheme } = useAIThemeGenerationCore();
  const isGeneratingTheme = false;
  const { getPreset } = useThemePresetStore();
  const currentPreset = themeState?.preset ? getPreset(themeState?.preset) : undefined;
  const isSavedPreset = !!currentPreset && currentPreset.source === "SAVED";

  const handleReset = () => {
    resetToCurrentPreset();
  };

  return (
    <div className="flex items-center gap-1">
      <MoreOptions disabled={isGeneratingTheme} />
      <Separator orientation="vertical" className="mx-1 h-8" />
      <ThemeToggle />
      <Separator orientation="vertical" className="mx-1 h-8" />
      <UndoRedoButtons disabled={isGeneratingTheme} />
      <Separator orientation="vertical" className="mx-1 h-8" />
      <ResetButton onClick={handleReset} disabled={!hasUnsavedChanges() || isGeneratingTheme} />
      <div className="hidden items-center gap-1 md:flex">
        <ImportButton onClick={onImportClick} disabled={isGeneratingTheme} />
      </div>
      <Separator orientation="vertical" className="mx-1 h-8" />
      {isSavedPreset && (
        <EditButton themeId={themeState.preset as string} disabled={isGeneratingTheme} />
      )}
      <ShareButton onClick={() => onShareClick(themeState.preset)} disabled={isGeneratingTheme} />
      <SaveButton onClick={onSaveClick} isSaving={isSaving} disabled={isGeneratingTheme} />
      <CodeButton onClick={onCodeClick} disabled={isGeneratingTheme} />
    </div>
  );
}
