<template>
  <Suspense>
    <template #default>
      <profile-layout v-if="hasGH" :github="github" />
    </template>
    <template #fallback>
      <div>Loading...</div>
    </template>
  </Suspense>
</template>

<script>
/* import Landing from './layouts/landing.vue'; */
import ProfileLayout from './layouts/profile.vue';

export default {

  components: {
    /* LandingLayout, */
    ProfileLayout
  },

  setup() {
    const [, user, repo] = window.location.pathname.split('/');
    return {
      hasGH: Boolean(user && repo),
      github: { user, repo }
    };
  }

};
</script>

