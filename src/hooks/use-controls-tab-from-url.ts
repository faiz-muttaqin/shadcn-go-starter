import { useEffect, useState } from "react";

const TABS = ["colors", "typography", "other", "ai"] as const;
export const DEFAULT_TAB = TABS[0];
export type ControlTab = (typeof TABS)[number];

/**
 * Lightweight client-side fallback for reading/writing the `tab` query param.
 * Uses `window.location` + `history.replaceState` so it doesn't require a
 * framework adapter (avoids nuqs adapter errors when running in plain
 * frontend environments).
 */
export const useControlsTabFromUrl = () => {
  const [tab, setTabState] = useState<ControlTab>(DEFAULT_TAB);

  useEffect(() => {
    if (typeof window === "undefined") return;
    const params = new URLSearchParams(window.location.search);
    const value = params.get("tab") ?? DEFAULT_TAB;
    if (!TABS.includes(value as ControlTab)) {
      console.warn(`Invalid tab value: ${value}. Falling back to default.`);
      // setTabState(DEFAULT_TAB);
      return;
    }
    // setTabState(value as ControlTab);
  }, []);

  const handleSetTab = (newTab: ControlTab) => {
    if (!TABS.includes(newTab)) {
      console.warn(`Invalid tab value: ${newTab}. Falling back to default.`);
      newTab = DEFAULT_TAB;
    }

    if (typeof window === "undefined") {
      setTabState(newTab);
      return;
    }

    const params = new URLSearchParams(window.location.search);
    params.set("tab", newTab);
    const newUrl = `${window.location.pathname}?${params.toString()}${window.location.hash}`;
    window.history.replaceState({}, "", newUrl);
    setTabState(newTab);
  };

  return { tab, handleSetTab };
};
