export default defineAppConfig({
  ui: {
    primary: "cool",
    input: {
      color: {
        white: {
          outline:
            "bg-white dark:bg-grey-900 text-grey-900 dark:text-white ring-1 ring-inset ring-gray-300 dark:ring-grey-0 focus:ring-2 focus:ring-primary-500 dark:focus:ring-primary-400",
        },
      },
    },
    select: {
      color: {
        white: {
          outline:
            "shadow-none bg-white dark:bg-grey-900 text-grey-900 ring-gray-300 dark:ring-grey-0",
        },
      },
    },
    button: {},
  },
});
