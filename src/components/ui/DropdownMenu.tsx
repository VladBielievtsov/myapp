import { cn } from "@/lib/utils";
import React, {
  useState,
  ReactNode,
  isValidElement,
  cloneElement,
  useEffect,
  useRef,
  forwardRef,
} from "react";

interface IRoot {
  children: ReactNode;
}

interface IToggleProps extends IRoot {
  toggleDropdown?: () => void;
}

interface IContentProps extends IRoot {
  isOpen?: boolean;
  align?: "right" | "left" | "center";
  className?: string;
}

export function DropdownMenu({ children }: IRoot) {
  const [isOpen, setIsOpen] = useState(false);

  const toggleDropdown = () => {
    setIsOpen(!isOpen);
  };

  const dropdownRef = useRef<HTMLDivElement>(null);
  const toggleRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node) &&
        toggleRef.current &&
        !toggleRef.current.contains(event.target as Node)
      ) {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  const clonedChildren = React.Children.map(children, (child) => {
    if (isValidElement(child)) {
      if (child.type === DropdownToggle) {
        return cloneElement(child, {
          toggleDropdown,
          ref: toggleRef,
        } as Partial<IToggleProps>);
      }
      if (child.type === DropdownContent) {
        return cloneElement(child, {
          isOpen,
          ref: dropdownRef,
        } as Partial<IContentProps>);
      }
    }
    return child;
  });

  return <div className="relative">{clonedChildren}</div>;
}

export const DropdownToggle = forwardRef<HTMLDivElement, IToggleProps>(
  ({ children, toggleDropdown }, ref) => {
    return (
      <div ref={ref} onClick={toggleDropdown}>
        {children}
      </div>
    );
  }
);

export const DropdownContent = forwardRef<HTMLDivElement, IContentProps>(
  ({ children, isOpen, align, className }, ref) => {
    return (
      isOpen && (
        <div
          ref={ref}
          className={cn(
            "absolute bg-white border border-zinc-200 rounded right-0 mt-1 overflow-hidden dark:bg-zinc-900 dark:border-zinc-600",
            align === "left" && "right-[unset] left-0",
            align === "right" && "left-[unset] right-0",
            align === "center" &&
              "right-[unset] left-1/2 transform -translate-x-1/2",
            className
          )}
        >
          {children}
        </div>
      )
    );
  }
);
