<template>
<v-container fluid fill-height>
  <v-row v-if="upcomingTasks.length !== 0" align="start" justify="center">
    <v-col cols="12">
      <div v-for="(task, index) in upcomingTasks" :key="task.id">
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
        You don't have any tasks upcoming!
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
  name: 'Upcoming',
  components: { Task },
  layout: 'tasky',
  computed: {
    ...mapGetters('task', ['upcomingTasks']),
  },
  mounted() {
    axios.get('/get_category/upcoming').then((resp) => {
      const tasks = JSON.parse(resp.data.tasks);
      this.setTasks({
        category: 'upcoming',
        tasks,
      });
    });
  },
  methods: {
    ...mapMutations('task', ['setTasks']),
  },
};
</script>
