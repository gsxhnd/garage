import { Directive, DirectiveBinding, watch } from "vue";
import {
  computePosition,
  flip,
  shift,
  offset,
  autoPlacement,
  platform,
  inline,
} from "@floating-ui/vue";
import { useElementHover } from "@vueuse/core";

export const tooltip: Directive = {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    const tooltipParent = document.createElement("div");
    const tooltipContent = document.createElement("div");
    const isHover = useElementHover(el);
    tooltipParent.className = "tooltip";
    tooltipContent.textContent = binding.value;
    tooltipContent.className = "tooltip-content";
    tooltipParent.appendChild(tooltipContent);

    watch(isHover, (newV, _oldV) => {
      if (newV) {
        document.body.insertAdjacentElement("beforeend", tooltipParent);
        computePosition(el, tooltipParent, {
          placement: "bottom",
          middleware: [offset(10), flip(), shift()],
          strategy: "absolute",
        }).then(({ x, y }) => {
          Object.assign(tooltipParent.style, {
            left: `${x}px`,
            top: `${y}px`,
          });
        });
        tooltipParent.style.display = "block";
      } else {
        tooltipParent.remove();
      }
    });
  },
};
