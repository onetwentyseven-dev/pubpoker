<template>
  <b-navbar toggleable="lg" type="light" fixed="top" class="transparent">
    <b-container>
      <b-navbar-brand to="/">
        <LogoLarge style="height: 100px" />
      </b-navbar-brand>
    </b-container>
  </b-navbar>
</template>

<style scoped>
.transparent {
  background: linear-gradient(
    rgba(255, 255, 255, 0.5),
    rgba(255, 255, 255, 0.8)
  );
  box-shadow: rgb(0 0 0 / 30%) 0px 0px 5px;
  backdrop-filter: blur(6px);
}
</style>
