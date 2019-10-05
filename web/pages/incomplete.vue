<template>
<v-container fluid fill-height>
  <v-row v-if="incompleteTasks.length !== 0" align="start" justify="center">
    <v-col cols="12">
      <div v-for="(task, index) in incompleteTasks" :key="task.id">
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
        You don't have any incomplete tasks!
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
  name: 'Incomplete',
  components: { Task },
  layout: 'tasky',
  computed: {
    ...mapGetters('task', ['incompleteTasks']),
  },
  mounted() {
    axios.get('/get_category/incomplete').then((resp) => {
      const tasks = JSON.parse(resp.data.tasks);
      this.setTasks({
        category: 'incomplete',
        tasks,
      });
    });
  },
  methods: {
    ...mapMutations('task', ['setTasks']),
  },
};
</script>
