<template>
  <div>
    <b-jumbotron class="p-0">
      <div class="banner">
        <h1 class="header">Charlotte's Best Free Texas Hold'em League</h1>
        <h4 class="header">Dealing Cards Since 2004</h4>
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
        <b-col lg="4">
          <b-card no-body>
            <template #header>
              <h4 class="mb-0">D.D. Peckers</h4>
            </template>
            <iframe
              src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3264.52435254214!2d-80.86941088394464!3d35.09361546957906!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x88569c59bc111a59%3A0xdffe8dbcf9ed2e33!2sD.D.%20PECKERS%20WING%20SHACK!5e0!3m2!1sen!2sus!4v1649819440638!5m2!1sen!2sus"
              height="300"
              loading="lazy"
            ></iframe>

            <b-list-group flush>
              <b-list-group-item
                ><b-icon icon="clock" class="mr-2" /> 8 p.m. on Mondays
              </b-list-group-item>
              <b-list-group-item href="tel:+17045414113">
                <b-icon icon="telephone" class="mr-2" />
                +1 (704) 541-4113
              </b-list-group-item>
              <b-list-group-item
                href="https://maps.google.com?q=10403 Park Rd, Charlotte, NC 28210"
                target="_blank"
              >
                <b-icon icon="map" class="mr-2" />10403 Park Rd, Charlotte, NC
                28210
              </b-list-group-item>
              <b-list-group-item href="https://ddpeckers.com" target="_blank">
                <b-icon icon="menu-app" class="mr-2" />
                View Menu
              </b-list-group-item>
            </b-list-group>
          </b-card>
        </b-col>
        <b-col lg="4">
          <b-card no-body>
            <template #header>
              <h4 class="mb-0">Zee's Tavern</h4>
            </template>
            <iframe
              src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d816.1313460921428!2d-80.86769827073005!3d35.09358969608172!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x88569c59bea5ab1f%3A0xd0ccbd3617a11f77!2sZee&#39;s%20Tavern!5e0!3m2!1sen!2sus!4v1649820853298!5m2!1sen!2sus"
              height="300"
              loading="lazy"
            ></iframe>
            <b-list-group flush>
              <b-list-group-item
                ><b-icon icon="clock" class="mr-2" /> 7 p.m. on Wednesdays
              </b-list-group-item>
              <b-list-group-item href="tel:+17045438862">
                <b-icon icon="telephone" class="mr-2" />
                +1 (704) 543-8862
              </b-list-group-item>
              <b-list-group-item
                href="https://maps.google.com?q=10403 Park Rd, Charlotte, NC 28210"
                target="_blank"
              >
                <b-icon icon="map" class="mr-2" />10403 Park Rd, Charlotte, NC
                28210
              </b-list-group-item>
              <b-list-group-item
                href="http://zees-tavern.edan.io/"
                target="_blank"
              >
                <b-icon icon="menu-app" class="mr-2" />
                View Amenities
              </b-list-group-item>
            </b-list-group>
          </b-card>
        </b-col>
        <b-col lg="4">
          <b-card no-body>
            <template #header>
              <h4 class="mb-0">Buffalo Wings And Rings</h4>
            </template>
            <iframe
              src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3267.2138736114985!2d-80.84103658401348!3d35.02637907321209!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x885682c3eb678017%3A0x5448f569497826ae!2sWings%20and%20Rings!5e0!3m2!1sen!2sus!4v1649820889425!5m2!1sen!2sus"
              height="300"
              loading="lazy"
            ></iframe>

            <b-list-group flush>
              <b-list-group-item
                ><b-icon icon="clock" class="mr-2" /> 7 p.m. on Thursdays
              </b-list-group-item>
              <b-list-group-item href="tel:+17045439464">
                <b-icon icon="telephone" class="mr-2" />
                +1 (704) 543-9464
              </b-list-group-item>
              <b-list-group-item
                href="https://maps.google.com?q=16715 Orchard Stone Run Charlotte, NC 28277"
                target="_blank"
              >
                <b-icon icon="map" class="mr-2" /> 16715 Orchard Stone Run
                Charlotte, NC 28277
              </b-list-group-item>
              <b-list-group-item
                href="https://www.wingsandrings.com/locations/charlotte"
                target="_blank"
              >
                <b-icon icon="menu-app" class="mr-2" />
                View Menu
              </b-list-group-item>
            </b-list-group>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<style scoped>
.banner {
  background: url('~assets/check-cards.jpg');
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
