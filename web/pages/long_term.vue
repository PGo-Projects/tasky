<template>
<v-container fluid fill-height>
  <v-row v-if="longTermTasks.length !== 0" align="start" justify="center">
    <v-col cols="12">
      <div v-for="(task, index) in longTermTasks" :key="task.id">
        <task
          :index="task.index"
          :name="task.name"
          :date="task.date"
          :time="task.time"
          :description="task.description"
          :category="task.category"
          :position="index"
          />
      </div>
    </v-col>
  </v-row>
  <v-row v-else align="center" justify="center">
    <v-col cols="12">
      <p class="display-3 text-center">
        You don't have any tasks long term!
      </p>
    </v-col>
  </v-row>
</v-container>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex';
import axios from 'axios';

import Task from '../components/Task.vue';

export default {
  name: 'LongTerm',
  components: { Task },
  layout: 'tasky',
  computed: {
    ...mapGetters('task', ['longTermTasks']),
  },
  mounted() {
    axios.get('/get_category/longTerm').then((resp) => {
      const tasks = JSON.parse(resp.data.tasks);
      this.setTasks({
        category: 'longTerm',
        tasks,
      });
    });
  },
  methods: {
    ...mapMutations('task', ['setTasks']),
  },
};
</script>
