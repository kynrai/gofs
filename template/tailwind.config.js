/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/ui/**/*.templ"],
  theme: {
    extend: {
      zIndex: {
        toast: 100,
      },
      keyframes: {
        zoomOut: {
          "0%": { transform: "scale(1)" },
          "100%": { transform: "scale(0.9)" },
        },
        bounceInDown: {
          "0%": { opacity: "0", transform: "translate3d(0, -3000px, 0)" },
          "60%": { opacity: "1", transform: "translate3d(0, 25px, 0)" },
          "75%": { transform: "translate3d(0, -10px, 0)" },
          "90%": { transform: "translate3d(0, 5px, 0)" },
          "100%": { transform: "none" },
        },
      },
    },
  },
  plugins: [],
};
