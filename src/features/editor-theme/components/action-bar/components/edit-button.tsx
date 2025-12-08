import { TooltipWrapper } from "@/components/TooltipWrapper";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { PenLine } from "lucide-react";
import { Link, useLocation, useSearch } from "@tanstack/react-router";

type EditButtonProps = React.ComponentProps<typeof Button> & {
  themeId: string;
};

export function EditButton({
  themeId,
  disabled,
  className,
  ...props
}: EditButtonProps) {
  const location = useLocation();
  const search = useSearch({ strict: false });

  const isEditing = location.pathname.includes(themeId);
  const href = `/editor/theme/${themeId}`;
  return (
    <TooltipWrapper label="Edit theme" asChild>
      <Link
        to={href}
        search={search} // ✅ Preserve current query params
        replace={isEditing} // optional: avoid extra history
      >
        <Button
          variant="ghost"
          size="sm"
          className={cn(className)}
          disabled={disabled || isEditing}
          {...props}
        >
          <PenLine className="size-3.5" />
          <span className="hidden text-sm md:block">Edit</span>
        </Button>
      </Link>
    </TooltipWrapper>
  );
}
