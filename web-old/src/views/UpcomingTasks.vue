<template>
<v-container fluid fill-height>
  <v-layout row align-start justify-center v-if="upcomingTasks.length !== 0">
    <v-flex xs12>
      <div v-for="(task, index) in upcomingTasks" :key="task.id">
        <task :index="task.index" :name="task.name" :date="task.date" :time="task.time"
              :description="task.description" :category="task.category" :position="index">
        </task>
      </div>
    </v-flex>
  </v-layout>
  <v-layout row align-center justify-center v-else>
    <v-flex xs12>
      <p class="display-3 text-xs-center">You don't have any upcoming tasks!</p>
    </v-flex>
  </v-layout>
</v-container>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex';
import axios from 'axios';

import Task from '@/components/Task.vue';

export default {
  components: { Task },
  name: 'upcomingtasks',
  mounted() {
    axios.get('/get_category/upcoming').then((resp) => {
      const tasks = JSON.parse(resp.data.tasks);
      this.setTasks({
        category: 'upcoming',
        tasks,
      });
    });
  },
  computed: {
    ...mapGetters('task', [
      'upcomingTasks',
    ]),
  },
  methods: {
    ...mapMutations('task', [
      'setTasks',
    ]),
  },
};
</script>
