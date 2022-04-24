<template>
  <Layout>
    <b-jumbotron class="p-0">
      <div class="banner text-center">
        <h1 class="header mx-3">Charlotte's Best Free Texas Hold'em League</h1>
        <h4 class="header mx-3">Dealing Cards Since 2004</h4>
      </div>
    </b-jumbotron>
    <b-container>
      <b-row>
        <b-col>
          <h2>Locations</h2>
          <hr />
        </b-col>
      </b-row>
      <b-row>
        <b-col
          md="6"
          sm="12"
          class="mb-2"
          v-for="location in locations"
          :key="location.name"
        >
          <b-card no-body>
            <template #header>
              <h4 class="mb-0">{{ location.name }}</h4>
            </template>
            <iframe
              :src="location.address.maps.iframe"
              height="300"
              loading="lazy"
            ></iframe>

            <b-list-group flush>
              <b-list-group-item
                ><b-icon icon="clock" class="mr-2" /> {{ location.hours }}
              </b-list-group-item>
              <b-list-group-item :href="'tel:' + location.telephone.raw">
                <b-icon icon="telephone" class="mr-2" />
                {{ location.telephone.formatted }}
              </b-list-group-item>
              <b-list-group-item
                :href="location.address.maps.link"
                target="_blank"
              >
                <b-icon icon="map" class="mr-2" />{{ location.address.text }}
              </b-list-group-item>
              <b-list-group-item :href="location.website.href" target="_blank">
                <b-icon icon="menu-app" class="mr-2" />
                {{ location.website.text }}
              </b-list-group-item>
            </b-list-group>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
  </Layout>
</template>
<script>
import Layout from "../layout/index.vue";

export default {
  components: {
    Layout,
  },
  data() {
    return {
      locations: [
        {
          name: "D.D. Peckers",
          hours: "8 p.m. on Mondays",
          telephone: {
            formatted: "+1 (704) 541-4113",
            raw: "+17045414113",
          },
          address: {
            maps: {
              iframe:
                "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3264.52435254214!2d-80.86941088394464!3d35.09361546957906!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x88569c59bc111a59%3A0xdffe8dbcf9ed2e33!2sD.D.%20PECKERS%20WING%20SHACK!5e0!3m2!1sen!2sus!4v1649819440638!5m2!1sen!2sus",
              link: "https://maps.google.com?q=10403 Park Rd, Charlotte, NC 28210",
            },
            text: "10403 Park Rd, Charlotte, NC 28210",
          },
          website: {
            href: "https://ddpeckers.com",
            text: "View Menu",
          },
        },
        {
          name: "Zee's Tavern",
          hours: "7 p.m. on Wednesdays",
          telephone: {
            formatted: "+1 (704) 543-8862",
            raw: "+17045438862",
          },
          address: {
            maps: {
              iframe:
                "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3264.525429036782!2d-80.86933978475722!3d35.09358858033556!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x88569c59bea5ab1f%3A0xd0ccbd3617a11f77!2sZee's%20Tavern!5e0!3m2!1sen!2sus!4v1650130315252!5m2!1sen!2sus",
              link: "https://maps.google.com?q=10403 Park Rd, Charlotte, NC 28210",
            },
            text: "10403 Park Rd, Charlotte, NC 28210",
          },
          website: {
            href: "http://zees-tavern.edan.io/",
            text: "View Amenities",
          },
        },
        {
          name: "Buffalo Wings And Rings",
          hours: "7 p.m. on Thursdays",
          telephone: {
            formatted: "+1 (704) 543-9464",
            raw: "+17045439464",
          },
          address: {
            maps: {
              iframe:
                "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3267.2138736114985!2d-80.84103658401348!3d35.02637907321209!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x885682c3eb678017%3A0x5448f569497826ae!2sWings%20and%20Rings!5e0!3m2!1sen!2sus!4v1649820889425!5m2!1sen!2sus",
              link: "https://maps.google.com?q=16715 Orchard Stone Run Charlotte, NC 28277",
            },
            text: "16715 Orchard Stone Run Charlotte, NC 28277",
          },
          website: {
            href: "https://www.wingsandrings.com/locations/charlotte",
            text: "View Menu",
          },
        },
      ],
    };
  },
  created() {
    console.log("Index");
  },
};
</script>
<style scoped>
.banner {
  background: url("../assets/check-cards.jpg");
  background-repeat: no-repeat;
  background-size: cover !important;
  height: 80vh;
  min-height: 700px;
  width: 100vw;
  background-position: top center;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  color: white;
}

.header {
  color: white;
  text-shadow: #000 1px 1px 10px;
}
</style>
