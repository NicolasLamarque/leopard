<!-- StatCard.vue -->
<template>
  <div
    class="bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm rounded-xl p-5 border border-gray-200 dark:border-gray-700 shadow-md hover:shadow-lg transition-all"
  >
    <div class="flex items-center justify-between">
      <div>
        <p class="text-sm text-gray-600 dark:text-gray-400 font-medium mb-1">
          {{ title }}
        </p>
        <p :class="['text-3xl font-bold', colorClasses.text]">
          {{ value }}
        </p>
      </div>
      <div
        :class="[
          'w-14 h-14 rounded-xl flex items-center justify-center shadow-md',
          colorClasses.bg,
        ]"
      >
        <component :is="iconComponent" :size="28" :class="colorClasses.icon" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import {
  PhoneIncoming,
  Clock,
  AlertCircle,
  CalendarCheck,
} from "lucide-vue-next";

const props = defineProps({
  title: String,
  value: [Number, String],
  icon: String,
  color: {
    type: String,
    default: "rose",
  },
});

const iconMap = {
  "phone-incoming": PhoneIncoming,
  clock: Clock,
  "alert-circle": AlertCircle,
  "calendar-check": CalendarCheck,
};

const iconComponent = computed(() => iconMap[props.icon] || PhoneIncoming);

const colorClasses = computed(() => {
  const colors = {
    rose: {
      bg: "bg-gradient-to-br from-rose-100 to-pink-100 dark:from-rose-900/30 dark:to-pink-900/30",
      text: "text-rose-600 dark:text-rose-400",
      icon: "text-rose-600 dark:text-rose-400",
    },
    amber: {
      bg: "bg-gradient-to-br from-amber-100 to-yellow-100 dark:from-amber-900/30 dark:to-yellow-900/30",
      text: "text-amber-600 dark:text-amber-400",
      icon: "text-amber-600 dark:text-amber-400",
    },
    red: {
      bg: "bg-gradient-to-br from-red-100 to-rose-100 dark:from-red-900/30 dark:to-rose-900/30",
      text: "text-red-600 dark:text-red-400",
      icon: "text-red-600 dark:text-red-400",
    },
    green: {
      bg: "bg-gradient-to-br from-green-100 to-emerald-100 dark:from-green-900/30 dark:to-emerald-900/30",
      text: "text-green-600 dark:text-green-400",
      icon: "text-green-600 dark:text-green-400",
    },
  };
  return colors[props.color] || colors.rose;
});
</script>
