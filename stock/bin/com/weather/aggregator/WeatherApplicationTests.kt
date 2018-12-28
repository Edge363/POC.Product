package com.weather.aggregator

import com.weather.aggregator.models.weather.WeatherCondition
import org.junit.Before
import org.junit.jupiter.api.Assertions.assertNotEquals
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.ValueSource
import org.junit.runner.RunWith
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.boot.test.web.client.TestRestTemplate
import org.springframework.test.context.junit4.SpringRunner

@RunWith(SpringRunner::class)
@SpringBootTest
class WeatherApplicationTests {

    private var client: TestRestTemplate = TestRestTemplate()

    @Before
    fun setup() {
    }

    @ParameterizedTest
    @ValueSource(ints = [524901, 30000])
    fun `Does service return something`(number: Int) {

        val result = client.getForEntity("http://localhost:8080/weather/$number", WeatherCondition::class.java)
        print("http://localhost:8080/weather/$number")
        assertNotEquals(result, null)

    }
}