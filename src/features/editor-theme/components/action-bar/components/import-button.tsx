import { TooltipWrapper } from "@/components/TooltipWrapper";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { FileCode } from "lucide-react";

type ImportButtonProps = React.ComponentProps<typeof Button>

export function ImportButton({ className, ...props }: ImportButtonProps) {
  return (
    <TooltipWrapper label="Import CSS variables" asChild>
      <Button variant="ghost" size="sm" className={cn(className)} {...props}>
        <FileCode className="size-3.5" />
        <span className="hidden text-sm md:block">Import</span>
      </Button>
    </TooltipWrapper>
  );
}
