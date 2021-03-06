<template>
<v-dialog persistent v-model="openTaskForm" max-width="1000px">
  <v-card>
    <v-container pa-12>
      <v-card-title>
        <span class="headline" :color="headerColor">{{ headerText }}</span>
      </v-card-title>

      <div class="my-4">
        <v-alert v-if="status !== undefined" :type="statusType"
                 value outlined>
          {{ status }}
        </v-alert>
      </div>
      
      <v-card-text>
        <v-form ref="form" @submit.prevent="saveForm">
          <v-row wrap>
            <v-col cols="12">
              <v-text-field
                v-model="info.name"
                :rules="nameRules"
                :value="info.name"
                :color="formColor"
                label="Name"
                name="name"
                ref="name"
                required
                />
            </v-col>
            
            <v-col cols="12" md="5">
              <v-menu
                ref="dateMenu"
                v-model="dateMenu"
                :color="formColor"
                :close-on-content-click="false"
                :return-value.sync="info.date"
                transition="scale-transition"
                offset-y
                min-width="280px"
                >
                <template v-slot:activator="{ on }">
                  <v-text-field
                    v-model="info.date"
                    :rules="dateRules"
                    :value="info.date"
                    :color="formColor"
                    label="Date"
                    readonly
                    @click="activateDatepicker"
                    @focus="dateMenu = true;">
                    <font-awesome-icon slot="prepend" :icon="['far', 'calendar-day']"
                                       style="font-size: 25px;" />
                  </v-text-field>
                </template>
                <v-date-picker v-model="info.date" :color="formColor"
                               :dark="useDarkTheme" no-title scrollable>
                  <v-spacer />
                  <v-btn text :color="formColor" @click="dateMenu = false;">
                    Cancel
                  </v-btn>
                  <v-btn text :color="formColor"
                         @click="$refs.dateMenu.save(info.date)">
                    OK
                  </v-btn>
                </v-date-picker>
              </v-menu>
            </v-col>
            <v-col md="1">
            </v-col>
            <v-col cols="12" md="5">
              <v-menu
                ref="timeMenu"
                v-model="timeMenu"
                :color="formColor"
                :close-on-content-click="false"
                :return-value.sync="info.time"
                transition="scale-transition"
                offset-y
                min-width="280px"
                @input="value"
                >
                <template v-slot:activator="{ on }">
                  <v-text-field
                    v-model="info.time"
                    :rules="timeRules"
                    :value="info.time"
                    :color="formColor"
                    label="Time"
                    prepend-icon="far-clock"
                    readonly
                    @click="activateTimepicker"
                    @focus="timeMenu = true;">
                    <font-awesome-icon slot="prepend" :icon="['far', 'clock']"
                                       style="font-size: 25px;" />
                    </v-text-field>
                </template>
                <v-time-picker v-if="taskOpenForm" v-model="info.time" ref="timePicker"
                               :color="formColor" :dark="useDarkTheme"
                               no-title scrollable>
                  <v-spacer />
                  <v-btn
                    text
                    :color="formColor"
                    @click="timeMenu = false;">
                    Cancel
                  </v-btn>
                  <v-btn
                    text
                    :color="formColor"
                    @click="$refs.timeMenu.save(info.time)">
                    OK
                  </v-btn>
                </v-time-picker>
              </v-menu>
            </v-col>
            
            <v-col cols="12">
              <v-textarea
                v-model="info.description"
                :value="info.description"
                :color="formColor"
                label="Description"
                name="description"
                rows=5
                auto-grow
                required
                />
            </v-col>
          </v-row>
          
          <v-row mx-2>
            <v-spacer />
            <v-btn type="submit" :color="formColor" text mr-5>
              {{ actionText }}
            </v-btn>
            <v-btn :color="formColor" text @click="closeTaskForm();">
              {{ cancelText }}
            </v-btn>
          </v-row>
        </v-form>
      </v-card-text>
    </v-container>
  </v-card>
</v-dialog>
</template>

<script>
import axios from 'axios';
import { mapActions, mapGetters, mapMutations } from 'vuex';

import taskMixin from '../mixins/task';

const NEW_ACTION = 'new';
const EDIT_ACTION = 'edit';

export default {
  name: 'TaskForm',
  mixins: [taskMixin],
  mounted() {
    this.$nextTick(() => this.$refs.name.focus());
  },
  computed: {
    ...mapGetters('status', [
      'status',
      'statusType',
    ]),
    ...mapGetters('taskForm', ['taskOpenForm']),
    actionText() {
      return (this.action === NEW_ACTION) ? 'Create' : 'Save';
    },
    actionLink() {
      return (this.action === NEW_ACTION) ? '/insert_task' : '/update_task';
    },
    cancelText() {
      return (this.action === NEW_ACTION) ? 'Close' : 'Cancel';
    },
    useDarkTheme() {
      return this.theme === 'dark';
    },
    formColor() {
      if (this.theme === 'light' && this.formcolor === undefined) {
        return 'info';
      }
      return this.formcolor;
    },
    headerColor() {
      if (this.headercolor === undefined) {
        return (this.theme === 'light') ? 'black' : 'white';
      }
      return this.headercolor;
    },
    headerText() {
      return (this.action === NEW_ACTION) ? 'Create Task' : 'Edit Task';
    },
    openTaskForm: {
      get() {
        return this.value;
      },
      set(val) {
        this.$emit('input', val);
      },
    },
  },
  data() {
    return {
      dateMenu: false,
      timeMenu: false,

      nameRules: [
        (v) => !!v || 'Name is required!',
      ],
      dateRules: [
        () => !!this.info.date || 'Date is required!',
      ],
      timeRules: [
        () => !!this.info.time || 'Time is required!',
      ],
    };
  },
  methods: {
    ...mapMutations('status', ['setError']),
    ...mapMutations('task', [
      'insertTask',
      'deleteTask',
      'updateTaskIndex',
    ]),
    ...mapActions('task', ['updateTask']),
    ...mapActions('taskForm', ['closeTaskForm']),
    activateDatepicker() {
      setTimeout(() => {
        this.dateMenu = true;
      }, 0);
    },
    activateTimepicker() {
      setTimeout(() => {
        this.timeMenu = true;
      }, 0);
    },
    async saveForm() {
      if (!this.$refs.form.validate()) {
        return;
      }

      const task = {
        index: this.info.index,
        name: this.info.name,
        date: this.info.date,
        time: this.info.time,
        description: this.info.description,
        position: this.info.position,

        oldCategory: this.info.category,
        oldPosition: this.info.position,
      };

      this.assignCategory(task);
      if (this.action === NEW_ACTION) {
        this.insertTask(task);
      } else {
        this.updateTask(task);
      }

      const resp = await fetch('/', response => response);
      axios.post(this.actionLink, task, {
        headers: { 'X-CSRF-TOKEN': resp.headers.get('X-CSRF-TOKEN') },
      }).then((response) => {
        if (response.data.statusType === 'success') {
          if (this.action === NEW_ACTION) {
            this.updateTaskIndex({
              category: task.category,
              position: task.position,
              index: parseInt(response.data.index, 10),
            });
          }
          this.closeTaskForm();
        } else {
          if (this.action === NEW_ACTION) {
            this.deleteTask(task);
          }
          this.setError(response.data.status);
        }
      }).catch((error) => {
        this.setError(error);
      });
    },
  },
  props: {
    info: undefined,
    action: {
      default: undefined,
      type: String,
      validator: value => [NEW_ACTION, EDIT_ACTION].includes(value),
    },

    // value is used to toggle the dialog
    value: {
      default: undefined,
    },
    theme: {
      default: 'light',
      type: String,
      validator: value => ['light', 'dark'].includes(value),
    },
    formcolor: {
      default: undefined,
      type: String,
    },
    headercolor: {
      default: undefined,
      type: String,
    },
  },
  watch: {
    openTaskForm() {
      this.$refs.form.resetValidation();
      this.$nextTick(() => this.$refs.name.focus());
    },
  },
};
</script>
