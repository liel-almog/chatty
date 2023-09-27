export const ENVIRONMENT_SETTINGS = {
    ['dev-' + import.meta.env['VITE_APP_NAME']]: 'פיתוח',
    ['stage-' + import.meta.env['VITE_APP_NAME']]: 'תצוגה',
    ['drills-' + import.meta.env['VITE_APP_NAME']]: 'תרגילית',
    [import.meta.env['VITE_APP_NAME']]: 'מבצעית',
};
